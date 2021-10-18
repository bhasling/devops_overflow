/***************
 * LinkCard.js
 * 
 * Implements a <LinkCard> that is a material design card used as a link to a nextjs route.
 ***************/
import {React, Component} from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Link from 'next/link';

class LinkCard extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             hover: false
         }
    }
 
    mouseEnter() {
        this.setState({'hover': true});
    }
    mouseLeave() {
        this.setState({'hover': false});
    }

    render() {
        return(
            <Link href={this.props.route ? this.props.route : ''} passHref>
                <Card style={{ 'width': '50%'}} 
                    elevation={this.state.hover? 12 : 3}
                    onMouseEnter={()=>this.mouseEnter()}
                    onMouseLeave={()=>this.mouseLeave()}
                    >
                    <CardContent>
                        <Typography fontWeight = 'bold' fontSize='1.5rem' color='primary'>
                            {this.props.name}
                        </Typography>
                        <Typography  color='primary'>
                            {this.props.summary}
                        </Typography>
                        <Typography  style={{'marginTop': '10px'}} color='primary' fontSize='0.8rem'>
                            {this.props.detail}
                        </Typography>
                    </CardContent>
                </Card>
            </Link>
        );
    }
}
export default LinkCard;

 