import { Outlet } from "react-router-dom";

import Sidebar from "../components/Sidebar";

function Layout() {
    return (
        <div className="flex min-h-screen">
            <Sidebar />

            <main className="grow p-8 bg-gray-50">
                <Outlet />
            </main>
        </div>
    );
}

export default Layout;
