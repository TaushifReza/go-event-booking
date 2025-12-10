function Dashboard() {
    return (
        <div className="space-y-6">
            <h1 className="text-3xl font-bold text-gray-900">
                Welcome to your Dashboard
            </h1>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div className="bg-white p-6 rounded-xl shadow-md">
                    <p className="text-sm font-medium text-gray-500">
                        Upcoming Events
                    </p>
                    <p className="text-2xl font-semibold text-gray-900 mt-1">
                        5
                    </p>
                </div>

                <div className="bg-white p-6 rounded-xl shadow-md">
                    <p className="text-sm font-medium text-gray-500">
                        Total Users
                    </p>
                    <p className="text-2xl font-semibold text-gray-900 mt-1">
                        1200
                    </p>
                </div>

                <div className="bg-white p-6 rounded-xl shadow-md">
                    <p className="text-sm font-medium text-gray-500">
                        Total Users
                    </p>
                    <p className="text-2xl font-semibold text-gray-900 mt-1">
                        1200
                    </p>
                </div>
            </div>

            <div className="mt-8 bg-white p-6 rounded-xl shadow-md">
                <h2 className="text-xl font-semibold mb-4">Activity Log</h2>
            </div>
        </div>
    );
}

export default Dashboard;
