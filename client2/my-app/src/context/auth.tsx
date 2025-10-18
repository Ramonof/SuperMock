import { BASE_URL } from '@/main'
import React, { createContext, useContext, useState, useEffect } from 'react'

interface User {
  id: string
  user: string
  email: string
}

interface AuthState {
  isAuthenticated: boolean
  user: User | null
  login: (username: string, password: string) => Promise<void>
  logout: () => void
}

const AuthContext = createContext<AuthState | undefined>(undefined)

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [isLoading, setIsLoading] = useState(true)

  // Restore auth state on app load
  useEffect(() => {
    const token = localStorage.getItem('auth-token')
    if (token) {
      // Validate token with your API
      fetch(BASE_URL + '/validate-token', {
        headers: { Authorization: `Bearer ${token}` },
      })
        .then((response) => response.json())
        .then((userData) => {
          if (userData.accessToken) {
            setUser(userData)
            setIsAuthenticated(true)
          } else {
            localStorage.removeItem('auth-token')
          }
        })
        .catch(() => {
          localStorage.removeItem('auth-token')
        })
        .finally(() => {
          setIsLoading(false)
        })
    } else {
      setIsLoading(false)
    }
  }, [])

  // Show loading state while checking auth
  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        Loading...
      </div>
    )
  }

  const login = async (user: string, pwd: string) => {
    // Replace with your authentication logic
    const response = await fetch(BASE_URL + '/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ user, pwd }),
    })

    if (response.ok) {
      const userData = await response.json()
      setUser(userData)
      setIsAuthenticated(true)
      // Store token for persistence
      localStorage.setItem('auth-token', userData.accessToken)
    } else {
      throw new Error('Authentication failed')
    }
  }

  const logout = () => {
    setUser(null)
    setIsAuthenticated(false)
    localStorage.removeItem('auth-token')
  }

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}