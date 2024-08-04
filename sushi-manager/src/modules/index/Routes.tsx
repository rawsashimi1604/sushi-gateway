import React from 'react'
import Header from '../../components/typography/Header'
import RouteDropdown from '../../components/sushi-gateway/RouteDropdown'

function Routes() {
    return (
        <div>
            <Header text="Routes" align="left" size="md" />
            <RouteDropdown />
        </div>
    )
}

export default Routes