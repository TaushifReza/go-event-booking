import { Link } from "react-router-dom";

function NotFound() {
    return (
        <main className="min-h-screen flex items-center justify-center bg-gray-50 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8 text-center p-10">
                <div className="mx-auto">
                    <p className="text-9xl font-extrabold text-blue-600">404</p>
                </div>

                <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
                    Page Not Found
                </h1>

                <p className="mt-4 text-lg text-gray-600">
                    Oops! The page you're looking for seems to have gone on a
                    little adventure. It might have been moved, deleted, or you
                    might have typed the address incorrectly.
                </p>

                <div className="mt-6 flex flex-col items-center justify-center space-y-4">
                    <Link
                        to="/"
                        className="w-full sm:w-auto px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition duration-150 ease-in-out"
                    >
                        Go to Homepage
                    </Link>
                </div>
            </div>
        </main>
    );
}

export default NotFound;
