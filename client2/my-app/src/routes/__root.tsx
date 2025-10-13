import Navbar from '@/components/generic/navbar'
import { Box, Flex, useColorModeValue, Text, Button } from '@chakra-ui/react'
import { createRootRouteWithContext, Link, Outlet, redirect, useLocation, useNavigate } from '@tanstack/react-router'
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

function RootLayout() {
  const { auth } = Route.useRouteContext()
  
  const location = useLocation()
  const navigate = useNavigate()
  function redirectOnClick(): void {
    navigate({
            to: '/login',
            search: {
              redirect: location.href,
            },
          })
  }

  return (
    <>
    <Navbar/>
      <Box bg={useColorModeValue("gray.400", "gray.700")} px={4} my={4} borderRadius={"5"}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
        <Flex
						justifyContent={"center"}
						alignItems={"center"}
						gap={3}
						display={{ base: "none", sm: "flex" }}
					>
        <Link to="/" className="[&.active]:font-bold">
          Home
        </Link>{' '}
        <Link to="/about" className="[&.active]:font-bold">
          About
        </Link>
        <Link to="/projects" className="[&.active]:font-bold">
          Projects
        </Link>
        </Flex>
        <Flex alignItems={"center"} gap={3}>
          {auth.isAuthenticated ? auth.user?.user : <Button onClick={redirectOnClick}>Sign In</Button>}
        </Flex>
        </Flex>
      </Box>
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  )
}

export const Route = createRootRouteWithContext<MyRouterContext>()({ 
  component: RootLayout,
})
