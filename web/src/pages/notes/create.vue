<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline">
          つれづれ日記
        </v-card-title>
        
        <v-card-text>
          <v-form
            ref="form"
            >
            <v-text-field
            v-model="postData.name"
            label="Name"
            required
            ></v-text-field>

            <v-text-field
            v-model="postData.message"
            label="Message"
            required
            ></v-text-field>

            <v-checkbox
            v-model="checkbox"
            :rules="[v => !!v || '投稿のためには同意してください。']"
            label="利用規約に同意して投稿する"
            required
            ></v-checkbox>

            <v-btn
            :disabled="submitted"
            color="success"
            class="mr-4"
            @click="onSubmit"
            >
            投稿する
            </v-btn>
            <span v-if="submitted" name="success-note-create">投稿に成功しました。</span>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/"> 一覧に戻る </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import axios from "axios"
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      submitted: false,
      checkbox: false,
      postData: {
        title: "",
        message: ""
      }
    }
  },
  methods: {
    onSubmit() : void {
      const res = axios.post(`http://localhost:8080/notes`, this.postData)
      .then((data) => {
        console.log(data)
        this.submitted = true
      })
      .catch((error) => {
        console.log(error)
      })
      .catch(() => {

      })

      return 
    }
  },
  components: {
  
  },
  async asyncData() {

  }
})
</script>
