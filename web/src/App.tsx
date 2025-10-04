import "./App.css";
import { useAuth } from "./contexts/AuthContext";

function App() {
  const { user } = useAuth();

  const handleInstallClick = () => {
    const state = user?.id || "user-session-placeholder";
    window.location.href = `https://github.com/apps/arcd-dev/installations/new?state=${state}`;
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-8">
      {user && (
        <div className="card">
          <h2 className="text-xl font-semibold mb-4">Welcome, {user.name}!</h2>
          <button
            onClick={handleInstallClick}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
          >
            Install GitHub App
          </button>
        </div>
      )}

      {!user && (
        <div className="text-center">
          <p className="text-gray-600">Please sign in to continue</p>
        </div>
      )}
    </div>
  );
}

export default App;
