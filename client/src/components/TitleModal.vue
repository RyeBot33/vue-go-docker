<template>
  <div class="text-center">
    <v-dialog
      v-model="dialog"
      width="500"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="red lighten-2"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Click Me
        </v-btn>
        <v-btn
          color="red lighten-2"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Click Me
        </v-btn>
      </template>

    <form>    
        <v-text-field
          v-model="name"
          :error-messages="nameErrors"
          label="Title Name"
        ></v-text-field>
    <v-btn class="mr-4" @click="addSideNav()">submit</v-btn>
    </form>
    </v-dialog>
  </div>
</template>>

<script>
export default {
    data() {
        return {
            dialog: false
        }
    },
}
</script>

<style>

</style>