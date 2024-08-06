import Header from "../../components/typography/Header";
import ServiceDropdown from "../../components/sushi-gateway/ServiceDropdown";

interface ServiceProps {
  data: any;
}

function Services({ data }: ServiceProps) {
  return (
    <div>
      <Header text="Services" align="left" size="md" />
      <div className="flex flex-col gap-3">
        {data?.map((service: any) => {
          return <ServiceDropdown key={service?.name} data={service} />;
        })}
      </div>
    </div>
  );
}

export default Services;
