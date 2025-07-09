const API_URL = "" 
export async function api_logout(){
    let response = await fetch(`${API_URL}/logout`)
    return response
}

export async function api_login(body){
    let response = await fetch(`${API_URL}/login`, {method: "POST", body: body})
    return response
}

export async function api_call(url, methodType="GET", body=""){
    let response = await fetch(`${API_URL}${url}`, {method: methodType, body: body})
    return response
}