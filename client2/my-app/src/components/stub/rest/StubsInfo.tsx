import { Badge, Box, Container, Divider, Flex, HStack, Link, Spinner, Stack, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Link as TanstackLink } from "@tanstack/react-router";
import { IoAdd, IoAddCircleOutline, IoMoon } from "react-icons/io5";
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
            </HStack>
            <Box
            maxW='fit-content' borderWidth='1px' borderRadius='lg' overflow='hidden'
                border={"1px"}
                borderColor={"gray.600"}
                p={3}
            >
                <Link as={TanstackLink}
                    to={`/project/${ProjectId}/rest/stubs/create`}
                    color={"yellow.100"}
                    variant="underline"
                    colorPalette="teal"
                >
                    <HStack>
                        <IoAddCircleOutline />
                        <Text>Stub</Text>
                    </HStack>
                </Link>
            </Box>
            <Divider p={1} width={'50%'} />
        </>
    )
};

export default StubsInfo;