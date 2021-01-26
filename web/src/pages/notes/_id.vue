<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline">
          つれづれ日記
        </v-card-title>
        
        <v-card-text>
          {{ note.message }}
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/notes/create"> 日記を書く </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">

import axios from "axios"
import Vue from "vue";
export default Vue.extend({
  components: {
  
  },
  async asyncData(params) {
    console.log(params.params.id)
    const res = await axios.get(`http://localhost:8080/notes/${params.params.id}`)
    const note = res.data.note
    console.log(note)
    return { note }
  }
})
</script>
