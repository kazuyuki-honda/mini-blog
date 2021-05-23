Vue.component('vueRemarkable', vueremarkable)


let showArticle= new Vue({
    el: "#show-article",
    filters: {
        formatDate(date) {
            return dateFns.format(new Date(date), 'YYYY-MM-DD HH:mm:ss');
        },
    },
    data:{
        articleID:null,
        article:'',
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
                this.article= response.data 
                console.log(this.article)
            })
            .catch(error =>{
                console.error(error.response)
            })
        },
        
    },

    created(){
        this.articleID=location.pathname.split('/')[3];
        console.log(this.articleID)
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

