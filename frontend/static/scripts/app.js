// import Vue from 'vue'

Vue.createApp({
  data() {
    return {
      fname: '',
      lname: '',
      idnumber: ''
    };
  },
  methods: {
    registerVoter() {
      const unique_hash = '';
      axios
        .post('http://localhost:8040/registrar' )
          .then(response => this.unique_hash = response.data)
          .catch(error => {
            this.errorMessage = error.message;
            console.error("There was an error!", error);
          });
      
          this.fname='';
          this.lname='';
          this.idnumber='';
    }
  }
}).mount("#app-1");


Vue.createApp({
  data() {
    return {
      selected: '',
      hashvalue: '',
      options: [
        { candidate: 'John', value: 'John', id: 1 },
        { candidate: 'Jacob', value: 'Jacob', id: 2},
        { candidate: 'Jeremy', value: 'Jeremy', id: 3},
        { candidate: 'Jedediah', value: 'Jedediah', id: 4},
      ],
    }
  },
  methods: {
    async submitVote() {

      const response = await axios.get("http://localhost:8040/vote/", this.hashvalue)
      
      //TODO: rewrite the POST request
      if (response == true){
        axios.post("http://localhost:8040/election_guard/vote/", this.option.id);
      }

    }
  }
}).mount("#app-2");


Vue.createApp({
  data() {
    return {
      verificationcode: '',
    };
  },
  methods: {
    async verify() {
      const response = await axios.get("http://localhost:8040/voters/", this.verificationcode)
    }

  }
}).mount("#app-3");