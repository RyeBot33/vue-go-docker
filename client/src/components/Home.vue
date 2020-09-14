<template>
  <div>
      <h1>{{msg}}</h1>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    mounted() {
        axios.get('/api').then(res => this.msg = res.data)
    },
    data() {
        return {
            msg: null
        }
    },
}
</script>

<style>

</style>