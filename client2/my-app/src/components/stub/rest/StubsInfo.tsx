import { Badge, Box, Container, Flex, HStack, Link, Spinner, Stack, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Link as TanstackLink } from "@tanstack/react-router";
import { IoAdd, IoMoon } from "react-icons/io5";
import { ExternalLinkIcon } from "lucide-react";

const StubsInfo = ({ ProjectId }: { ProjectId: string }) => {

    return (
        <>
            <HStack spacing='24px' justify={"center"}>
                <Text
                    fontSize={"4xl"}
                    textTransform={"uppercase"}
                    fontWeight={"bold"}
                    textAlign={"center"}
                    my={2}
                    bgGradient='linear(to-l, #0bf827ff, #4000ffff)'
                    bgClip='text'
                >
                    Stubs
                </Text>
                <Flex
                    // flex={0.8}
                    alignItems={"center"}
                    border={"1px"}
                    borderColor={"gray.600"}
                    p={2}
                    borderRadius={"lg"}
                    justifyContent={"space-between"}
                >
                    <Link as={TanstackLink}
                        to={`/project/${ProjectId}/rest/stubs/create`}
                        color={"yellow.100"}
                        variant="underline"
                        colorPalette="teal"
                    >
                        <HStack>
                            <IoAdd />
                            <Text>Stub</Text>
                        </HStack>
                    </Link>
                </Flex>
            </HStack>
        </>
    )
};

export default StubsInfo;