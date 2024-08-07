import Header from "../../components/typography/Header";
import RouteDropdown from "../../components/sushi-gateway/RouteDropdown";
import NormalText from "../../components/typography/Text";

interface RoutesProps {
  data: any;
}

function Routes({ data }: RoutesProps) {
  return (
    <div>
      <Header text="Routes" align="left" size="md" />
      {(data?.length === 0 || data === undefined) && (
        <NormalText text="No routes were found." />
      )}
      <div className="flex flex-col gap-3">
        {data?.map((route: any) => {
          return <RouteDropdown data={route} />;
        })}
      </div>
    </div>
  );
}

export default Routes;
