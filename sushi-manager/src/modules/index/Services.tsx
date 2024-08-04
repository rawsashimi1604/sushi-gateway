import Header from '../../components/typography/Header'
import ServiceDropdown from '../../components/sushi-gateway/ServiceDropdown'

function Services() {
    return (
        <div>
            <Header text="Services" align="left" size="md" />
            <ServiceDropdown />

        </div>
    )
}

export default Services