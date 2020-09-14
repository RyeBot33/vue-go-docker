<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>
        <h1>{{ msg }}</h1>
      <v-col class="mb-4">
     </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
  export default {
    name: 'HelloWorld',
    mounted() {
      axios.get('/api').then(res => this.msg = res.data)
    },
    data: () => ({
      msg: null
    }),
  }
</script>
