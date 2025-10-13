import Navbar from '@/components/generic/navbar'
import { createRootRouteWithContext, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'

interface AuthState {
  isAuthenticated: boolean
  user: { id: string; user: string; email: string } | null
  login: (username: string, password: string) => Promise<void>
  logout: () => void
}

interface MyRouterContext {
  auth: AuthState
}

const RootLayout = () => (
  <>
  <Navbar/>
    <div className="p-2 flex gap-2">
      <Link to="/" className="[&.active]:font-bold">
        Home
      </Link>{' '}
      <Link to="/about" className="[&.active]:font-bold">
        About
      </Link>
      <Link to="/projects" className="[&.active]:font-bold">
        Projects
      </Link>
    </div>
    <hr />
    <Outlet />
    <TanStackRouterDevtools />
  </>
)

export const Route = createRootRouteWithContext<MyRouterContext>()({ 
  component: RootLayout,
})
