<template>
  <form>
    <v-text-field
      v-model="path"
      :counter="10"
      label="Path"
      required
      @input="$v.name.$touch()"
      @blur="$v.name.$touch()"
    ></v-text-field>
    <v-text-field
      v-model="name"
      label="Name"
      required
      @input="$v.email.$touch()"
      @blur="$v.email.$touch()"
    ></v-text-field>
     <v-text-field
      v-model="component"
      label="Component"
      required
      @input="$v.email.$touch()"
      @blur="$v.email.$touch()"
    ></v-text-field>
    <v-text-field
      v-model="icon"
      label="Icon"
      required
      @input="$v.email.$touch()"
      @blur="$v.email.$touch()"
    ></v-text-field>
    <v-checkbox
      v-model="subgroups"
      label="Has subgroups?"
    ></v-checkbox>

    <v-btn class="mr-4" @click="addSideNav()">submit</v-btn>
   
  </form>
</template>>

<script>
import axios from 'axios'
export default {
    mounted() {
        axios.get('/api').then(res => this.msg = res.data)
    },
    data() {
        return {
            msg: null,
            path: null,
            name: null,
            component: null,
            icon: null,
            subgroups: false,
            person: {
                FirstName: "Ryan",
                LastName: "Brothers",
                Email: "Test@test.com",
                Password: "hello"
            }
        }
    },
    methods: {
        addSideNav(){
            console.log(JSON.stringify(this.person))
            axios.post('/api/write', JSON.stringify(this.person)).then(res => console.log(res.data))
        }
    },
}
</script>

<style>

</style>