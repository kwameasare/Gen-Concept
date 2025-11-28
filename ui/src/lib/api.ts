const BASE_URL = "http://localhost:5005/api/v1";

export const setToken = (token: string) => localStorage.setItem("token", token);
export const getToken = () => localStorage.getItem("token");
export const removeToken = () => localStorage.removeItem("token");

export async function request<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    const url = `${BASE_URL}${endpoint}`;
    const token = getToken();

    const headers: HeadersInit = {
        "Content-Type": "application/json",
        ...options.headers,
    };

    if (token) {
        (headers as any)["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(url, {
        ...options,
        headers,
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error?.message || error.message || "Something went wrong");
    }

    const json = await response.json();
    return json.result;
}

export const api = {
    post: <T>(endpoint: string, body: any) =>
        request<T>(endpoint, {
            method: "POST",
            body: JSON.stringify(body),
        }),
    get: <T>(endpoint: string) => request<T>(endpoint, { method: "GET" }),
    put: <T>(endpoint: string, body: any) =>
        request<T>(endpoint, {
            method: "PUT",
            body: JSON.stringify(body),
        }),
    delete: <T>(endpoint: string) => request<T>(endpoint, { method: "DELETE" }),
};
