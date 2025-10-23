import Navbar from '@/components/generic/navbar'
import { Box, Flex, useColorModeValue, Text, Button, useColorMode } from '@chakra-ui/react'
import { createRootRouteWithContext, Link, Outlet, redirect, useLocation, useNavigate } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'
import { IoMoon } from 'react-icons/io5'
import { LuSun } from 'react-icons/lu'

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
  
  const { colorMode, toggleColorMode } = useColorMode();
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
    {/* <Navbar/> */}
      <Box bg={useColorModeValue("gray.400", "gray.700")} px={4} my={4} borderRadius={"5"} width={'100%'}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
        <Flex
						justifyContent={"center"}
						alignItems={"center"}
						gap={3}
						display={{ base: "none", sm: "flex" }}
					>
        <Link to="/" className="[&.active]:font-bold">
          SuperMock
        </Link>{' '}
        <Link to="/about" className="[&.active]:font-bold">
          About
        </Link>
        <Link to="/projects" className="[&.active]:font-bold">
          Projects
        </Link>
        </Flex>
        <Flex alignItems={"center"} gap={3}>
          {/* Toggle Color Mode */}
          <Button onClick={toggleColorMode}>
            {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
          </Button>
          {auth.isAuthenticated ? <Flex gap={3} h={16} alignItems={"center"} justifyContent={"space-between"}><Text> {auth.user?.user} </Text><Button onClick={auth.logout}>Sign Out</Button></Flex> : <Button onClick={redirectOnClick}>Sign In</Button>}
        </Flex>
        </Flex>
      </Box>
      {/* <hr /> */}
      <Outlet />
      {/* <TanStackRouterDevtools /> */}
    </>
  )
}

export const Route = createRootRouteWithContext<MyRouterContext>()({ 
  component: RootLayout,
})
