import { HStack, Stack, Box, Text, useColorModeValue, Grid, GridItem, Link } from '@chakra-ui/react'
import { createFileRoute, Outlet } from '@tanstack/react-router'
import { Link as TanstackLink } from "@tanstack/react-router";

export const Route = createFileRoute('/about')({
  component: About,
})

function About() {
  const { auth } = Route.useRouteContext()
  const color = useColorModeValue("gray.800", "yellow.100")
  
  return <div className="p-2">
  <Grid
  templateAreas={`"header header"
                  "nav main"
                  "nav footer"`}
  gridTemplateRows={'50px 1fr 30px'}
  gridTemplateColumns={'150px 1fr'}
  h='200px'
  gap='1'
  color='blackAlpha.700'
  fontWeight='bold'
>
  <GridItem pl='2' area={'nav'}>
    <Link as={TanstackLink}
      to={`/about/overview`}
    >
      <Text color={color}>Обзор</Text>
    </Link>{" "}
    <Link as={TanstackLink}
      to={`/about/comparison`}
    >
      <Text color={color}>Сравнение с WireMock OSS</Text>
    </Link>{" "}
  </GridItem>
  <GridItem pl='2' area={'main'}>
    <Outlet/>
  </GridItem>
</Grid>
  </div>
}