// import Vue from 'vue'
// import qs from 'qs'

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
      const data =  {
        first_name: this.fname,
        last_name: this.lname,
        idnumber: this.idnumber
      }

      // const params = new URLSearchParams();
      // params.append( 'first_name', 'Cooper');
      // params.append( 'last_name', 'this.lname');
      // params.append( 'id', 'this.idnumber');

      console.log(data);
      // console.log(params);

      // axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';

      axios.post('http://localhost:8040/registrar',
          // headers: { 'content-type': 'application/json' },
          data,
          // data: qss.stringify(sdata) ,
          // headers: {
          //   'Accept': 'application/json',
          //   'content-type': 'application/x-www-form-urlencoded',
          // },
              
        )
        .then(function (response) {
          console.log(response);
        })
        .catch(error => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });

        this.fname = '';
        this.lname = '';
        this.idnumber = '';
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
      axios
      .get('http://localhost:8040/registrar', {
        params: {
          verificationcode: this.verificationcode
        }
      })
        .then(response => console.log(response))
        .catch(error => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
      // const response = await axios.get("http://localhost:8040/registrar", this.verificationcode)
      this.verificationcode = '';
    }

  }
}).mount("#app-3");