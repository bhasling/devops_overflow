/***************
 * Detail Page to view an issue after searching.
 * 
 ***************/
import {React, Component} from 'react';
import Button from '@mui/material/Button';
import Icon from '@mui/material/Icon';
import IconButton from '@mui/material/IconButton';

class DetailSection extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             errorMessage: ""
         }
    }

    clickSearch() {
        this.props.setIssue(null);
    }

    render() {
        const formContainerStyle = {
            'display': 'flex',
            'columnGap': '30px',
            'margin': '20px'
        }
        const formEntryStyle = {
            'width': '50%'
        }
        const fullLineStyle = {
            'width': '100%',
        }
        return (
            <div>
                <center>
                <IconButton style={{'verticalAlign': 'top', 'marginRight': '10px'}} 
                            onClick={()=>this.clickSearch()}
                            color='primary'>
                    <Icon style={{'fontSize':'2rem'}}>search</Icon>
                </IconButton>
                </center>

                <div style={formContainerStyle}>
                        <div>{this.props.issue.title}</div>
                        <div>{this.props.issue.product}</div>
                </div>
                <div style={formContainerStyle}>
                    <div style={fullLineStyle}>
                        {this.props.issue.description}
                    </div>
                </div>

            </div>
        )

    }
}
export default DetailSection;

