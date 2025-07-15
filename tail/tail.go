package tail

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// FileState represents the current state of a file
type FileState struct {
	Size    int64
	ModTime time.Time
	Inode   uint64
}

// OutputWriter defines the interface for writing output
type OutputWriter interface {
	Write(data []byte) (int, error)
}

// File represents a file being monitored
type File struct {
	Path      string
	Handle    *os.File
	Output    OutputWriter
	State     FileState
	InitialState FileState
}

// getFileState retrieves the current state of a file
func getFileState(file *os.File) (FileState, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return FileState{}, err
	}

	// Get inode information
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return FileState{}, fmt.Errorf("failed to get file system info")
	}

	return FileState{
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime(),
		Inode:   stat.Ino,
	}, nil
}

// updateFileState updates the current file state
func (f *File) updateFileState() error {
	state, err := getFileState(f.Handle)
	if err != nil {
		return err
	}
	f.State = state
	return nil
}

// checkFileState compares current file state with initial state to detect changes
func (f *File) checkFileState() error {
	// First check if the file still exists at the filesystem level
	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		return fmt.Errorf("file %s no longer exists", f.Path)
	}

	currentState, err := getFileState(f.Handle)
	if err != nil {
		// If we can't get file state from handle, check if it's because file was deleted/replaced
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s no longer exists", f.Path)
		}
		return fmt.Errorf("failed to get current file state: %w", err)
	}

	// Check if file was deleted or replaced (inode changed)
	if currentState.Inode != f.InitialState.Inode {
		return fmt.Errorf("file %s was renamed or replaced (inode changed from %d to %d)", 
			f.Path, f.InitialState.Inode, currentState.Inode)
	}

	// Check if file was truncated (size decreased)
	if currentState.Size < f.State.Size {
		return fmt.Errorf("file %s was truncated (size decreased from %d to %d)", 
			f.Path, f.State.Size, currentState.Size)
	}

	// Update current state
	f.State = currentState
	return nil
}

// OpenFile validates and opens a file for monitoring
func OpenFile(filePath string) (*File, error) {
	return OpenFileWithOutput(filePath, os.Stdout)
}

// OpenFileWithOutput validates and opens a file for monitoring with custom output
func OpenFileWithOutput(filePath string, output OutputWriter) (*File, error) {
	// Validate file path is not empty
	if filePath == "" {
		return nil, fmt.Errorf("file path cannot be empty")
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	// Open the file in read-only mode
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	// Get initial file state
	initialState, err := getFileState(file)
	if err != nil {
		return nil, fmt.Errorf("failed to get initial file state for %s: %w", filePath, err)
	}

	return &File{
		Path:      filePath,
		Handle:    file,
		Output:    output,
		State:     initialState,
		InitialState: initialState,
	}, nil
}

// SeekToEnd positions the file pointer at the end of the file
func (f *File) SeekToEnd() error {
	_, err := f.Handle.Seek(0, 2) // 2 = io.SeekEnd
	if err != nil {
		return fmt.Errorf("failed to seek to end of file %s: %w", f.Path, err)
	}
	return nil
}

// Close closes the file handle
func (f *File) Close() error {
	if f.Handle != nil {
		err := f.Handle.Close()
		f.Handle = nil // Set to nil after closing to prevent double-close issues
		return err
	}
	return nil
}

// Monitor continuously reads and outputs new lines from the file
func (f *File) Monitor(stopChan <-chan struct{}) error {
	buffer := make([]byte, 4096)
	lineBuffer := ""
	lastCheckTime := time.Now()
	checkInterval := 100 * time.Millisecond // Check file state every 100ms for faster detection
	
	for {
		select {
		case <-stopChan:
			// Flush any remaining incomplete line when stopping
			if lineBuffer != "" {
				f.Output.Write([]byte(lineBuffer))
			}
			return nil
		default:
			// Check file state periodically for log rotation
			if time.Since(lastCheckTime) >= checkInterval {
				if err := f.checkFileState(); err != nil {
					return err
				}
				lastCheckTime = time.Now()
			}
			
			// Read available data
			n, err := f.Handle.Read(buffer)
			if err != nil {
				// Check if it's just EOF (no more data available)
				if err.Error() == "EOF" {
					// Check if we have any incomplete line in buffer
					if lineBuffer != "" {
						f.Output.Write([]byte(lineBuffer))
						lineBuffer = ""
					}
					time.Sleep(100 * time.Millisecond)
					continue
				}
				// Check for other temporary errors that might occur during file operations
				if os.IsNotExist(err) {
					return fmt.Errorf("file %s no longer exists: %w", f.Path, err)
				}
				return fmt.Errorf("error reading file %s: %w", f.Path, err)
			}
			
			if n == 0 {
				// No data read, wait a bit
				time.Sleep(100 * time.Millisecond)
				continue
			}
			
			// Process the read data
			data := string(buffer[:n])
			lineBuffer += data
			
			// Process complete lines
			for {
				newlineIndex := -1
				for i, char := range lineBuffer {
					if char == '\n' {
						newlineIndex = i
						break
					}
				}
				
				if newlineIndex == -1 {
					// No complete line found, keep the data in buffer
					break
				}
				
				// Extract and print the complete line
				line := lineBuffer[:newlineIndex+1]
				f.Output.Write([]byte(line))
				
				// Remove the processed line from buffer
				lineBuffer = lineBuffer[newlineIndex+1:]
			}
		}
	}
}
