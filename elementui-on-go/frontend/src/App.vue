<template>
  <div id="app">
    <el-container>
      <el-container>
        <el-aside width="200px">
          <el-col :span="24">
            <el-menu default-active="2" class="el-menu-vertical-demo" :default-openeds="['2', '3']">
              <el-menu-item index="1">Sample Corporation</el-menu-item>
              <el-submenu index="2">
                <template slot="title"><i class="el-icon-menu"></i>営業管理</template>
                <el-menu-item index="2-1">> 案件一覧</el-menu-item>
                <el-menu-item index="2-2">> 案件登録</el-menu-item>
              </el-submenu>
              <el-submenu index="3">
                <template slot="title"><i class="el-icon-setting"></i>システム設定</template>
                <el-menu-item index="3-1" disabled>> 社員登録申請</el-menu-item>
                <el-menu-item index="3-2">> 社員名簿
                </el-menu-item>
                <el-menu-item index="3-3">> 設定</el-menu-item>
              </el-submenu>
            </el-menu>
          </el-col>
        </el-aside>
        <el-container>
          <el-header style="text-align: right;">
            <el-dropdown>
              <i class="el-icon-setting" style="margin-right: 15px font-size: 30px;"></i>
              <i class="el-icon-caret-bottom" style="margin-right: 15px font-size: 30px;"></i>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>ログアウト</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <span></span>
          </el-header>
          <el-main>
            <el-table 
              :data="staffs" 
              style="width: 100%"
              :row-class-name="roleDisplay">
              <el-table-column prop="id" label="社員ID"></el-table-column>
              <el-table-column prop="name" label="氏名"></el-table-column>
              <el-table-column prop="email" label="メールアドレス"></el-table-column>
              <el-table-column prop="role" label="権限" width="180"></el-table-column>
            </el-table>
            <el-pagination
              background
              layout="prev, pager, next"
              :total="50">
             </el-pagination>
          </el-main>
        </el-container>
      </el-container>
      <el-footer>Copy right reserverd by khigsahiguchi.com</el-footer>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'app',
  methods: {
    roleDisplay({row}) {
      if (row.role ==='システム管理者') {
        return 'admin-row';
      } else if (row.role == 'マネージャー') {
        return 'manager-row';
      } else {
        return '';
      }
    }
  },
  data() {
    return {
      staffs : [{
        id: 1,
        name: '鈴木太郎',
        email: 'example1@example.com',
        role: 'システム管理者'
      }, {
        id: 2,
        name: '山田花子',
        email: 'example2@example.com',
        role: 'システム管理者'
      }, {
        id: 3,
        name: '岡本太郎',
        email: 'example3@example.com',
        role: 'マネージャー'
      }, {
        id: 4,
        name: '矢野一郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 5,
        name: '矢野二郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 6,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 7,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 8,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 9,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 10,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 11,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }, {
        id: 12,
        name: '鈴木太郎',
        email: 'example4@example.com',
        role: '一般社員'
      }]
    }
  }
}
</script>

<style>
#app {
  font-family: Helvetica, sans-serif;
  text-align: center;
}
.el-table .admin-row {
  background: oldlace
}
.el-table .manager-row {
  background: #f0f9eb
}
.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}
</style>
