import { NavLink } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

function PageNav() {
    const { isAuthenticated } = useAuth();
    const baseLink =
        "text-gray-600 uppercase text-lg font-semibold hover:text-blue-600 transition";
    const activeLink = "text-blue-600";

    const ctaLink =
        "bg-blue-600 text-white px-4 py-2 rounded-lg font-semibold uppercase text-lg shadow hover:bg-blue-700 transition";

    return (
        <nav className="flex items-center justify-between py-4 px-6">
            <ul className="flex items-center gap-16">
                <li>
                    <NavLink
                        to="/dashboard"
                        className={({ isActive }) =>
                            isActive ? `${baseLink} ${activeLink}` : baseLink
                        }
                    >
                        Dashboard
                    </NavLink>
                </li>

                <li>
                    {isAuthenticated ? (
                        ""
                    ) : (
                        <NavLink
                            to="/login"
                            className={({ isActive }) =>
                                isActive ? ctaLink : ctaLink
                            }
                        >
                            Login
                        </NavLink>
                    )}
                </li>
            </ul>
        </nav>
    );
}

export default PageNav;
