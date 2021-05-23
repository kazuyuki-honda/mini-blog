let editArticle= new Vue({
    el: "#edit-article",

    filters: {
        formatDate(date) {
            return dateFns.format(new Date(date), 'YYYY-MM-DD HH:mm:ss');
        },
    },
    data:{
        articleID:null,
        // article:'',
        requestBody:{
            title:'',
            body:'',
        },
        errMessage:[]
    },
    methods:{
        search(){
            const url='/api/articles/'+this.articleID
            const headers = {
                'Content-Type': 'application/json; charset=UTF-8'
            }
           axios
            .get(url,{headers:headers,data:{}
            })
            .then(response =>{
                this.requestBody= response.data 
            })
            .catch(error =>{
                console.error(error.response)
            })
        },
        
        edit(){
            // console.log(this.requestBody)
            if( this.articleID !=null){
                var url='/api/articles/'+this.articleID
            }else{
                var url ='/api/articles/new'
            }
            const headers = {
                'Content-Type': 'application/json; charset=UTF-8'
            }

            const postData = this.requestBody
           axios
            .post(url,postData,{headers:headers})
            .then(response =>{
                alert("登録完了しました");
                window.location.href = "/admin";
                // console.log(this.article)
            })
            .catch(error =>{
                console.error(error.response.data)
                let message = error.response.data.message
                if (Array.isArray(message)){
                    this.errMessage=message
                }else{
                this.errMessage.push(message)
                }
            })
        }

    },

    created(){
        this.articleID=location.pathname.split('/')[4];
        if (this.articleID !=null){
        this.search();
        }
    }
})


const logout = new Vue({
    el: "#logout",

    methods:{
    logout() {
        localStorage.removeItem('token')
        location.href = '/login'
    },
    }
})