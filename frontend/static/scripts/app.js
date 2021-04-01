// import Vue from 'vue'

Vue.createApp({
  data() {
    return {

    };
  },
  methods: {
    registerVoter() {
      
    }
  }
}).mount("#app-1");


Vue.createApp({
  data() {
    return {
      selected: '',
      options: [
        { candidate: 'John', value: 'John'},
        { candidate: 'Jacob', value: 'Jacob'},
        { candidate: 'Jeremy', value: 'Jeremy'},
        { candidate: 'Jedediah', value: 'Jedediah'},
      ]
    }
  },
  methods: {

  }
}).mount("#app-2");


Vue.createApp({
  data() {
    return {

    };
  },
  methods: {

  }
}).mount("#app-3");