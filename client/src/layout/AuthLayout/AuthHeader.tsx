import Container from "@/components/Container";
import { Title, Text } from "@/components/Typo";
import { Link, useMatch } from "react-router-dom";
export default function AuthHeader() {
  const isAtLogin = useMatch("login");
  return (
    <Container>
      <div className="py-4 px-4 flex justify-between items-baseline">
        <Title as="h2">Audit Log</Title>
        <div className="flex space-x-2">
          <Text>or</Text>
          <Link className="underline" to={isAtLogin ? "/register" : "/login"}>
            {isAtLogin ? "Register" : "Login"}
          </Link>
        </div>
      </div>
    </Container>
  );
}
