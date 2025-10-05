import { Box, Container, Flex, Link, Text } from "@chakra-ui/react";
import { useColorModeValue } from "./ui/color-mode";

export default function SideBar() {
    return (
        <Container maxW={"900px"}>
            <Box bg={useColorModeValue("gray.400", "gray.700")} px={4} my={4} borderRadius={"5"}>
                <Flex
                    justifyContent={"center"}
                    alignItems={"center"}
                    gap={3}
                    display={{ base: "none", sm: "flex" }}
                >
                    <Link
                                        color={"yellow.100"}
                                        textDecoration={"none"}
                                        variant="underline"
                                        href={`/projects`}
                                        colorPalette="teal"
                                    >
                                        Projects
                                    </Link>{" "}
                </Flex>
            </Box>
        </Container>
    )
}
