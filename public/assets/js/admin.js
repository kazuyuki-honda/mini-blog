Vue.component('paginate', VuejsPaginate)


let articleList= new Vue({
    el: "#article-list",

    filters: {
        formatDate(date) {
            return dateFns.format(new Date(date), 'YYYY-MM-DD HH:mm:ss');
        },
    },
    data:{
        allArticles:"",
        requestBody: {
        limit: 5, // 取得件数
        page: 1,
        },
        currentPage: 1,
        pages: null
    },

    methods:{
        deleteArticle(id){
            console.log(id)
        },

        paging(page) {
            this.requestBody = {
                ...this.requestBody,
                page,
            }
            this.search();
        },

        search(){
            // this.requestBody ={
            //     ...this.requestBody,
            //     page:1,
            // }
            let postData = this.requestBody
            const url='/api/articles/list'
            const headers = {
                'Content-Type': 'application/json; charset=UTF-8'
            }
           axios
            .post(url,postData, {headers:headers})
            .then(response =>{
                const { articles, pages, page, total} = response.data;
                this.allArticles = articles;
                this.pages = pages;
                this.currentPage = page;
                this.total = total;
                // this.allArticles = response.data
                console.log(this.allArticles)
            })
            .catch(error =>{
                console.error(error.response)
            })
        }
    
    },

    created(){
        this.search();
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