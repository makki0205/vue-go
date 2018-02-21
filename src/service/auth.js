import axios from 'axios'

export default auth;
var auth = new Auth()
class Auth{
    constructor(){
        this.token = ""
    }
    Login :async (email, password) => {
        const hoge = await axios.post("/api/signin")
        console.log(hoge)

    }
    SetToken(t){
        this.token = t
    }
}
