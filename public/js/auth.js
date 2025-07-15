const API_URL = "" 
export async function api_logout(){
    return api_call(`${API_URL}/logout`, "POST")
}

export async function api_login(body){
    return api_call(`${API_URL}/login`, "POST", body)
}

export async function api_get(url){
    let response = await fetch(`${API_URL}${url}`)
    return response
}

export async function api_call(url, methodType, body=""){
    let response = await fetch(`${API_URL}${url}`, {method: methodType, body: body})
    return response
}
