import Header from "../../components/typography/Header";
import ServiceDropdown from "../../components/sushi-gateway/ServiceDropdown";
import NormalText from "../../components/typography/Text";

interface ServiceProps {
  data: any;
}

function Services({ data }: ServiceProps) {
  return (
    <div>
      <Header text="Services" align="left" size="md" />
      {(data?.length === 0 || data === undefined) && (
        <NormalText text="No services were found." />
      )}
      <div className="flex flex-col gap-3">
        {data?.map((service: any) => {
          return <ServiceDropdown key={service?.name} data={service} />;
        })}
      </div>
    </div>
  );
}

export default Services;
