import { useColorModeValue } from "@chakra-ui/react"

export function getColor() {
    const color = useColorModeValue("gray.800", "yellow.100")
    return color
}

export default getColor;