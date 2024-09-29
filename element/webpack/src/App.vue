<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import { useRoute } from 'vue-router';

export default {
  name: 'App',
  async mounted() {
    const { query } = useRoute();
    await this.loadData(query);
  },
  methods: {
      /**
     * 初始化
     */
     async loadData(options) {
      let username = options?.username
			if (username != null && username.length > 0) {
				localStorage.setItem("user-token", username)
			}
    },

  },
}
</script>

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
