import { Link } from "react-router-dom";

function Register() {
    return (
        <main className="min-h-screen flex flex-col items-center justify-center bg-gray-50 px-4">
            <section className="w-full max-w-md bg-white rounded-xl shadow-lg p-8 mt-10">
                <h1 className="text-2xl font-bold text-gray-900 text-center">
                    Create Your Account
                </h1>
                <p className="mt-2 text-center text-gray-600">
                    Sign up to start booking and managing your events.
                </p>

                <form className="mt-6 space-y-4">
                    <div>
                        <label
                            className="block text-gray-700 font-semibold mb-1"
                            htmlFor="name"
                        >
                            Full Name
                        </label>
                        <input
                            type="text"
                            id="name"
                            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-600 focus:outline-none"
                            placeholder="John Doe"
                        />
                    </div>

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
                        />
                    </div>

                    <div>
                        <label
                            className="block text-gray-700 font-semibold mb-1"
                            htmlFor="confirmPassword"
                        >
                            Confirm Password
                        </label>
                        <input
                            type="password"
                            id="confirmPassword"
                            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-600 focus:outline-none"
                            placeholder="Confirm your password"
                        />
                    </div>

                    <button
                        type="submit"
                        className="w-full mt-4 px-4 py-2 bg-blue-600 text-white font-semibold rounded-lg shadow hover:bg-blue-700 transition"
                    >
                        Register
                    </button>
                </form>

                <p className="mt-6 text-center text-gray-600">
                    Already have an account?{" "}
                    <Link
                        to="/login"
                        className="text-blue-600 font-semibold hover:underline"
                    >
                        Login
                    </Link>
                </p>
            </section>
        </main>
    );
}

export default Register;
