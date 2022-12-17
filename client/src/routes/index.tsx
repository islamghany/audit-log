import { useRoutes } from "react-router-dom";
import routes from "./initializeRoutes";

const Routes = () => {
  const elements = useRoutes(routes);

  return elements;
};

export default Routes;
