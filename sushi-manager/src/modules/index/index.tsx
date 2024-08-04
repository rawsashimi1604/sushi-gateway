import Container from "../../components/layout/Container";
import Global from "./Global";
import Json from "./Json";
import Routes from "./Routes";
import Services from "./Services";


function IndexModule() {

  // Get some information from Sushi proxy API

  return (
    <Container>
      <div className="flex flex-col gap-2">
        <Global />
        <Services />
        <Routes />
        <Json />
      </div>
    </Container>
  );
}

export default IndexModule;
