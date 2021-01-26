<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline">
          つれづれ日記
        </v-card-title>
        
        <v-card-text>
          <div class="" v-for="note in notes" :key="note.id">
            <nuxt-link :to="{name:'notes-id', params: {id: note.id}}">{{ note.message }}</nuxt-link>
            
          </div>
          
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
  async asyncData() {
    const res = await axios.get(`http://localhost:8080/notes`)
    const notes = res.data.notes
    return { notes }
  }
})
</script>
