<script lang="ts" setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { Atest } from '../../wailsjs/go/Demo/Demo'
import { defineComponent } from 'vue';
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  console.log("1111")
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

function demo() {
  Atest(data.name).then(result =>{
    data.resultText = result
  })
}

</script>

<template>

  <a-space direction="vertical">
    <div id="result" class="result">{{ data.resultText }}</div>
    <a-input v-model:value.lazy="data.name" autofocus placeholder="Lazy usage" />
    <a-button type="primary" @click="greet">确认</a-button>
    <a-button type="primary" @click="demo">确认demo</a-button>
  </a-space>
</template>

<style scoped>

</style>
