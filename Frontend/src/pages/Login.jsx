import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { useAuth } from "../context/AuthContext";

function Login() {
    const [email, setEmail] = useState("user1@gmail.com");
    const [password, setPassword] = useState("Admin123@");
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const { login, isLoading } = useAuth();

    async function handleSubmit(e) {
        e.preventDefault();

        if (!email || !password) {
            setError("Please fill email and password");
            return;
        }

        const result = await login(email, password);

        if (result.success) {
            navigate("/");
        } else {
            setError(result.error || "An unexpected login error occurred");
        }
    }

    const loading = isLoading;

    return (
        <main className="min-h-screen flex flex-col items-center justify-center bg-gray-50 px-4">
            <section className="w-full max-w-md bg-white rounded-xl shadow-lg p-8 mt-10">
                <h1 className="text-2xl font-bold text-gray-900 text-center">
                    Login
                </h1>
                <p className="mt-2 text-center text-gray-600">
                    Enter your credentials to access your events.
                </p>

                <form className="mt-6 space-y-4" onSubmit={handleSubmit}>
                    <div>
                        <label
                            className="block text-gray-700 font-semibold mb-1"
                            htmlFor="email"
                        >
                            Email
                        </label>
                        <input
                            type="email"
                            id="email"
                            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-600 focus:outline-none"
                            placeholder="you@example.com"
                            required
                            onChange={(e) => setEmail(e.target.value)}
                            value={email}
                        />
                    </div>

                    <div>
                        <label
                            className="block text-gray-700 font-semibold mb-1"
                            htmlFor="password"
                        >
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-600 focus:outline-none"
                            placeholder="Enter your password"
                            required
                            onChange={(e) => setPassword(e.target.value)}
                            value={password}
                        />
                    </div>

                    {error && (
                        <div className="text-red-600 text-center font-medium p-2 bg-red-50 rounded-lg">
                            {error}
                        </div>
                    )}

                    <button
                        type="submit"
                        className="w-full mt-4 px-4 py-2 bg-blue-600 text-white font-semibold rounded-lg shadow hover:bg-blue-700 transition"
                        disabled={loading}
                    >
                        {loading ? "Logging.." : "Login"}
                    </button>
                </form>

                <p className="mt-6 text-center text-gray-600">
                    Don't have an account?{" "}
                    <Link
                        to="/register"
                        className="text-blue-600 font-semibold hover:underline"
                    >
                        Register
                    </Link>
                </p>
            </section>
        </main>
    );
}

export default Login;
