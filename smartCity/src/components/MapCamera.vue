<script setup lang="ts">
import { OrbitCamera, HTML, Find } from "lingo3d-vue"
import { reactive } from 'vue';

//global state that determines the selected object
//定义选中物体的全局状态
import objectSelectedState from "../states/objectSelectedState"

const props = defineProps({
  name: {
    type: String,
    required: true
  },
  title: {
    type: String,
    required: true
  }
})
const host = 'http://localhost'
const buildingMessages = reactive({
  Height: '',
  Coordinate: '',
  curPower: '',
  PeopleNum: ''
})
interface keyTrans {
  [name: string]: string
}
const buildingKey: keyTrans = {
  [`01-shanghaizhongxindasha`]: 'shanghaitower',
  [`02-huanqiujinrongzhongxin_huanqiujinrongzhongxin_0`]: 'shanghaiwfc',
  [`03-jinmaodasha_jjinmaodasha_0`]: 'jinmaotower',
  [`04-dongfangmingzhu_dongfangmingzhu_0`]: 'orientalpearltower'
}

const firstPost = (name: string) => {
  const url = `${host}/${name}`
  fetch(`${url}/height`,{
    mode:"no-cors"
  })
    .then(response => response.text())
    .then(json => {
      console.log(json)
      buildingMessages.Height = json
    })
  fetch(`${url}/coordinate`,{
    mode:"no-cors"
  })
    .then(response => response.text())
    .then(json => {
      console.log(json)
      buildingMessages.Coordinate = json
    })
  fetch(`${url}/cur_power`,{
    mode:"no-cors"
  })
    .then(response => response.text())
    .then(json => {
      console.log(json)
      buildingMessages.curPower = json
    })
  fetch(`${url}/num_people`,{
    mode:"no-cors"
  })
    .then(response => response.text())
    .then(json => {
      console.log(json)
      buildingMessages.PeopleNum = json
    })
}
let timer: number | undefined
const steadyPost = (name: string) => {
  const url = `${host}/${name}`
  timer = window.setInterval(() => {
    if (objectSelectedState.key !== name) {
      window.clearInterval(timer)
      timer = undefined
    }
    fetch(`${url}/cur_power`,{
    mode:"no-cors"
  })
      .then(response => response.text())
      .then(json => {
        console.log(json)
        buildingMessages.curPower = json
      })
    fetch(`${url}/num_people`,{
    mode:"no-cors"
  })
      .then(response => response.text())
      .then(json => {
        console.log(json)
        buildingMessages.PeopleNum = json
      })
    console.log(timer);
  }, 2000)
}

//callback that gets called when the user clicks on an object
//当用户点击物体时的回调函数
const handleClick = (e: Object) => {
  if (timer) {
    window.clearInterval(timer)
    timer = undefined
  }
  console.log(e, props.name);
  if (objectSelectedState.name === props.name) {
    objectSelectedState.name = ""
  }
  else {
    const key=buildingKey[props.name]
    objectSelectedState.name = props.name
    objectSelectedState.key=key
    firstPost(key)
    steadyPost(key)
  }
  console.log(objectSelectedState.name);
}
</script>
  
<template>
  <!-- find object by name, highlight outline if selected -->
  <!-- 通过名称来查询物体，选中时轮廓高亮 -->
  <Find :name="props.name"
        :id="props.name"
        :outline="objectSelectedState.name === props.name"
        @click="handleClick">
    <!-- HTML tag that follows object -->
    <!-- 跟随物体的HTML标签 -->
    <HTML>
    <div class="p-2 backdrop-blur-xl absolute text-white pointer-events-auto cursor-pointer rounded-md overflow-hidden spAdd"
         @mousedown="handleClick">
      {{ title }}
    </div>
    <div class="messageCard p-2 backdrop-blur-xl absolute text-white pointer-events-auto cursor-pointer rounded-md overflow-hidden spAdd"
         v-if="objectSelectedState.name === props.name">

      <div>Height: {{buildingMessages.Height}}</div>
      <div>Coordinate: {{buildingMessages.Coordinate}}</div>
      <div>Current Power: {{buildingMessages.curPower}}</div>
      <div>Current Number of People: {{buildingMessages.PeopleNum}}</div>

    </div>

    </HTML>
  </Find>
  <!-- Camera that activates when object is selected -->
  <!-- 当物体被选中时启动的相机 -->
  <OrbitCamera :fov="90"
               :targetId="name"
               :active="objectSelectedState.name === name"
               :transition="0.02"
               :innerZ="50"
               :innerY="-30"
               enableDamping
               :minPolarAngle="120"
               :maxPolarAngle="120"
               autoRotate />
</template>
<style>
.messageCard {
  position: absolute;
  left: 100px;
  min-width: 300px;
}
</style>
  