Vue.createApp({
  data() {
    return {
      candidates: [],
      enteredValue: ''
    };
  },
  methods: {
    addCandidate() {
      this.candidates.push(this.enteredValue);
      this.enteredValue = '';
    }
  }
}).mount('#app');
