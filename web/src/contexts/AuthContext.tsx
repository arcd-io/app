import { createContext, useContext, useState, useEffect, ReactNode } from "react";
import { authClient, getSessionToken, setSessionToken, clearSessionToken } from "../lib/auth";
import type { User } from "../gen/auth/v1/auth_pb";

interface AuthContextType {
  user: User | null;
  loading: boolean;
  login: () => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    checkSession();
  }, []);

  async function checkSession() {
    const token = getSessionToken();
    if (!token) {
      setLoading(false);
      return;
    }

    try {
      const response = await authClient.getSession({ token });
      setUser(response.user ?? null);
    } catch (error) {
      console.error("Session validation failed:", error);
      clearSessionToken();
      setUser(null);
    } finally {
      setLoading(false);
    }
  }

  function login() {
    window.location.href = "http://localhost:8080/api/auth/github";
  }

  function logout() {
    clearSessionToken();
    setUser(null);
  }

  return (
    <AuthContext.Provider value={{ user, loading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
