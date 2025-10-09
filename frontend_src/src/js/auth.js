const API_URL = ( import.meta.env.DEV ? import.meta.env.VITE_API_URL : "")

export async function api_logout(){
    return api_call(`/logout`, "POST")
}

export async function api_login(body){
    return api_call(`/login`, "POST", body)
}

export async function api_get(url){
    try{
        let response = await fetch(`${API_URL}${url}`, {credentials: "include"})
        return response
    } catch(error) {
        console.error(error.message)
        throw(error)
    }
    
}

export async function api_call(url, methodType, body=""){
    try{
        let response = await fetch(`${API_URL}${url}`, {method: methodType, body: body, credentials: "include"})
        return response
    } catch(error) {
        console.error(error.message)
        throw(error)
    }
}

export async function getMealsInfo(){
    try{
        let response = await api_get("/api/settings/all")
        if(response.ok){
            let data = await response.json()
            return JSON.parse(data.data)
        }
    } catch(error){
        console.error(error.message)
        throw(error)
    }
}