import { Link } from "react-router-dom";
import PageNav from "../components/PageNav";
import { useAuth } from "../context/AuthContext";

function HomePage() {
    const { isAuthenticated } = useAuth();

    return (
        <main className="min-h-screen flex flex-col items-center bg-gray-50">
            <PageNav />

            <section className="max-w-3xl text-center mt-20 px-4">
                <h1 className="text-4xl md:text-5xl font-bold leading-tight text-gray-900">
                    Book Your Events Seamlessly.
                    <br />
                    <span className="text-blue-600">EventMaster</span> helps you
                    organize and manage events effortlessly.
                </h1>

                <h2 className="mt-6 text-lg md:text-xl text-gray-700 leading-relaxed">
                    From concerts and workshops to conferences and meetups,
                    track, plan, and invite attendees all in one place. Never
                    miss an important event, and make sure every experience is
                    memorable.
                </h2>

                {isAuthenticated ? (
                    ""
                ) : (
                    <Link
                        to="/login"
                        className="inline-block mt-10 px-8 py-3 text-lg font-semibold rounded-lg shadow
                    bg-blue-600 text-white hover:bg-blue-700 transition"
                    >
                        Start Booking Now
                    </Link>
                )}
            </section>
        </main>
    );
}

export default HomePage;
