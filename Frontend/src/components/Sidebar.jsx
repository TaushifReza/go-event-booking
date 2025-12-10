// src/components/Sidebar.jsx

import React from "react";
import { NavLink } from "react-router-dom";

function Sidebar() {
    // Define your navigation links
    const navLinks = [{ name: "Dashboard", path: "/dashboard" }];

    return (
        <aside className="w-64 bg-gray-800 text-white min-h-screen p-4 flex flex-col">
            <div className="text-2xl font-bold mb-8 text-blue-400">
                App Name
            </div>

            <nav className="space-y-2">
                {navLinks.map((link) => (
                    <NavLink
                        key={link.name}
                        to={link.path}
                        // Tailwind classes for styling NavLink
                        className={({ isActive }) =>
                            `flex items-center p-3 rounded-lg transition duration-200 ${
                                isActive
                                    ? "bg-blue-600 text-white shadow-lg"
                                    : "hover:bg-gray-700 text-gray-300"
                            }`
                        }
                    >
                        {link.name}
                    </NavLink>
                ))}
            </nav>
        </aside>
    );
}

export default Sidebar;
