import { Stack, Container } from "@chakra-ui/react";
import { Toaster } from "../components/ui/toaster";

const Main = () => {
    return (
    <Stack h="100vh">
      <Container>
        <Toaster />
      </Container>
    </Stack>
  )
}

export default Main;