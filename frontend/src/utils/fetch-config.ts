export const BACKENDURL = "http://localhost:8080";

export const ENDPOINTS = {
    users: "/users",
    userByID: (id: number) => `/users/${id}`,
    updateCartByID: (id: number) => `/cart/${id}`,
    productsByPage: (page: number) => `/products/page/${page}`,
    productByCode: (code: string) => ("/products/code/"+code)
}

export function SetRequestConfig(method: ("GET" | "POST" | "PUT" | "DELETE"), body?: BodyInit): RequestInit {
    return {
        headers: {
            "Content-Type": "application/json"
        },
        method,
        body
    }
}