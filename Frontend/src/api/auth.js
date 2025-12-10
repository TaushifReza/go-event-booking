const API_BASE_URL = import.meta.env.VITE_API_URL;

export async function loginUser(email, password) {
    try {
        const res = await fetch(`${API_BASE_URL}/auth/login/`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });

        const data = await res.json();

        if (data.success) {
            return {
                success: data.success,
                message: data.message,
                data: data.data,
            };
        } else {
            return {
                success: data.success,
                message: data.message,
                error: data.error,
            };
        }
    } catch (error) {
        console.error(`ERROR: ${error}`);
        return { success: false, message: error.message };
    }
}
