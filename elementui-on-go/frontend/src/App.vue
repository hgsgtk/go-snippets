<template>
  <div id="app">
    <div class="block">
      <span class="slider">Slider</span>
      <el-slider v-model="ruleForm.age"></el-slider>
    </div>
    <div>{{ ruleForm.age }}</div>

    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm">
      <el-form-item label="年齢" prop="age">
        <el-input type="number" v-model="ruleForm.age" auto-complete="off"></el-input>
      </el-form-item>
    </el-form>

    <span class="datepicker">日付</span>
    <el-date-picker
      v-model="value2"
      type="date"
      placeholder="choice date">
    </el-date-picker>
  </div>
</template>

<script>
export default {
  name: 'app',
  data() {
    return {
      ruleForm: {
        age: 0,
      },
      rules: {
        age: [
          {
            validator: (rule, value, callback) => {
              if (!value) {
                return callback(new Error('値を入力してください'));
              }
              if (value > 51) {
                return callback(new Error('年齢は50歳以下にしてください'))
              }
            },
            trigger: 'change'
          }
        ]
      },
      value2: '',
    }
  }
}
</script>

<style>
#app {
  font-family: Helvetica, sans-serif;
  text-align: center;
}
</style>
