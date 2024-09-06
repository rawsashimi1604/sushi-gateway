import { IoMdInformationCircle } from "react-icons/io";
import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";

function RoutesModule() {
  return (
    <Container>
      <DashboardCard>
        <div className="p-6">
          <Header text="routes  " align="left" size="sm" />
          <div>
            <table className="w-full text-sm text-left rtl:text-right">
              <thead className="text-xs uppercase">
                <tr className="font-lora font-light tracking-widest">
                  <th className="pl-0 px-6 py-3">
                    <div className="flex flex-row items-center gap-2">
                      <span>name</span>
                      <IoMdInformationCircle className="text-lg mb-0.5" />
                    </div>
                  </th>
                  <th className="px-6 py-3">
                    <div className="flex flex-row items-center gap-2">
                      <span>path</span>
                      <IoMdInformationCircle className="text-lg mb-0.5" />
                    </div>
                  </th>
                  <th className="px-6 py-3">
                    <div className="flex flex-row items-center gap-2">
                      <span>methods</span>
                      <IoMdInformationCircle className="text-lg mb-0.5" />
                    </div>
                  </th>
                  <th className="px-6 py-3">
                    <div className="flex flex-row items-center gap-2">
                      <span>service</span>
                      <IoMdInformationCircle className="text-lg mb-0.5" />
                    </div>
                  </th>
                </tr>
              </thead>
              <tbody className="font-lora tracking-wider">
                <tr className="border-b">
                  <td
                    scope="row"
                    className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
                  >
                    sushi-route
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    /sushi
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    GET
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    SushiService
                  </td>
                </tr>
                <tr className="border-b">
                  <td
                    scope="row"
                    className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
                  >
                    sushi-route
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    /sushi
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    GET
                  </td>

                  <td
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap"
                  >
                    SushiService
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </DashboardCard>
    </Container>
  );
}

export default RoutesModule;
