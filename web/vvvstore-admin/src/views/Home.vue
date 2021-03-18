<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
    <a-button type="primary" @click="showMessage">
      Primary
    </a-button>
  </div>
</template>

<script>
import { onMounted } from 'vue';
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld
  },
  emits: ['click'],
  setup() {

    onMounted(() => {
      console.log('Component is mounted!');
    });

  
  },
  methods: {
    showMessage() {
      this.$message.info('this is info message');

      // 监听事件
      this.$event.on('foo', msg => console.log(msg));

      // 发送事件
      this.$event.emit('foo', 'send foo event');

      // 移除事件
      this.$event.off('foo', () => console.log('remove foo event'));  
    }
  },
};
</script>
