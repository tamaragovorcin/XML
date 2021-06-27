import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_AGENT } from "../constants.js";
import Axios from "axios";
import { Carousel } from 'react-responsive-carousel';
import { AiFillDelete } from 'react-icons/ai';
import ImageUploader from 'react-images-upload';
import Order from "../components/Order";
import { GiLargeDress } from "react-icons/gi";
class ProfilePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            albums: [],
            key: 1 | props.activeKey,
            hiddenEditInfo: true,
            price: "",
            quantity: "",
            name: "",
            id: "",
            user: "",
            product: "",
            showOrderModal: false,
            name: "",
            price: "",
            pictures: [],
            openModal: false,
            help: [],
            fileUploadOngoing: false,
            albumId: "",


        }

        this.handleSelect = this.handleSelect.bind(this);
        this.onDrop = this.onDrop.bind(this);
    }



    delete = (e, name, id) => {
        let help = this.state.albums
        
        console.log(help)
        let ime = ""
        let helpMedia = []
        let helpMediaOrig = []


        help.forEach(album => {
                let i = 0
                helpMedia = album.Media
                helpMediaOrig = album.MediaOrig
            album.Media.forEach(post => {
                if (post === name) {
                    ime = album.MediaOrig[i]
                    album.MediaOrig.splice(album.MediaOrig[i], 1);
                    album.Media.splice(album.Media[i], 1);
                }
                i = i + 1
                 

            })
        })

        console.log(help)
      


        this.setState({
            albums: help
        });

        let deleteDTO = {
            image: ime,
            AlbumId: id,
        }

        Axios.post(BASE_URL_AGENT + "/api/removeImage", deleteDTO)
            .then((res) => {

                console.log(res.data)
                alert("success")
                window.location.reload();

            })
            .catch((err) => {
                console.log(err);
            });


    }
    onDrop(picture) {
        this.setState({
            pictures: [],
        });
        this.setState({
            pictures: this.state.pictures.concat(picture),
        });

    }
    addImage(e, postId) {
        let h = []
		this.state.pictures.forEach((pic) => {
			h.push(pic.name)
			this.setState({
				help: this.state.help.concat(pic.name),
			});
		});

        const product = {
			postId: postId,
			media: h
		
		};
		this.sendRequestForFeedAlbum(product);


    }







    sendRequestForFeedAlbum(feedPostDTO) {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		Axios.post(BASE_URL_AGENT + "/api/addImages" , feedPostDTO)
			.then((res) => {
			
					//this.props.openModal=true 
					//this.props.redirect=true
				
				let feedId = res.data;

				this.state.pictures.forEach((pic) => {
					this.test(pic, id, feedId);
				});

				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added album feed post." });
                window.location.reload();

			})
			.catch((err) => {
				console.log(err);
			});
	}




    test(pic, userIdd, feedId) {
        this.setState({
            fileUploadOngoing: true
        });

        const fileInput = document.querySelector("#fileInput");
        const formData = new FormData();

        formData.append("file", pic);
        formData.append("test", "StringValueTest");

        const options = {
            method: "POST",
            body: formData

        };
        fetch(BASE_URL_AGENT + "/api/image/" + userIdd, options);
    }

    AddToCart = (id) => {
        this.setState({ product: id });
        this.setState({ showOrderModal: true });

    };
    handleSelect(key) {
        this.setState({ key })
    }
    handleNameChange = (event, id) => {
        let help = this.state.albums
        help.forEach(post => {
            if (post.Id === id) {
                post.Name = event.target.value
            }
        })

        this.setState({ albums: help });
    };

    handlePriceChange = (event, id) => {
        let help = this.state.albums
        help.forEach(post => {
            if (post.Id === id) {
                post.Price = event.target.value
            }
        })

        this.setState({ albums: help });
    };
    handleOrderModalClose = () => {
        this.setState({ showOrderModal: false });
    };
    handleQuantityChange = (event, id) => {
        let help = this.state.albums
        help.forEach(post => {
            if (post.Id === id) {
                post.Quantity = event.target.value
            }
        })

        this.setState({ albums: help });
    };

    handleChangeAlbum = (event) => {
        this.setState({ hiddenEditInfo: false });
    };

    removeProduct = (idd) => {

        Axios.get(BASE_URL_AGENT + "/api/product/remove/" + idd)
            .then((res) => {

                console.log(res.data)
                alert("success")
                window.location.reload();

            })
            .catch((err) => {
                console.log(err);
            });

    }

    SendChange = (idd) => {

        let help = []

        this.state.albums.forEach(a => {
            const product = {
                user: a.User,
                name: a.Name,
                price: a.Price,
                quantity: a.Quantity,

            };

            help.push(product)
        })







        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);
        Axios.post(BASE_URL_AGENT + "/api/feedAlbum/edit/" + id, this.state.albums)
            .then((res) => {

                console.log(res.data)
                alert("success")

            })
            .catch((err) => {
                console.log(err);
            });
    }




    componentDidMount() {

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);
        Axios.get(BASE_URL_AGENT + "/api/feedAlbum/all/" + id)
            .then((res) => {
                this.setState({ albums: res.data });
                console.log(res.data)
             
            })
            .catch((err) => {
                console.log(err);
            });
    }

    render() {
        return (
            <React.Fragment>
                <TopBar />
                <Header />


                <div className="d-flex align-items-center" style={{ marginLeft: "25rem", marginTop: "10rem" }} >
                    <div className="container-fluid">

                        <table className="table">
                            <tbody>
                                {this.state.albums.map((post) => (

                                    <tr id={post.id} key={post.id}>

                                        <tr style={{ width: "100%" }}>
                                            <td colSpan="3"  style={{ width: "45rem" }}>
                                                <Carousel dynamicHeight={true}>
                                                    {post.Media.map(img => (<div>
                                                        <img
                                                            className="img-fluid"
                                                            src={`data:image/jpg;base64,${img}`}
                                                            width="100%"
                                                            alt="description"
                                                        />
                                                       
                                                    </div>))}
                                                </Carousel>
                                            </td>
                                        </tr>
                                        
                                  
                                        <tr>
                                            <td colSpan="3">
                                                <div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
                                                    <label>Product name</label>
                                                    <br></br>
                                                    <input style={{ width: "45rem" }}
                                                        readOnly={this.state.hiddenEditInfo}
                                                        className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
                                                        placeholder="Product name"
                                                        type="text"
                                                        onChange={e => this.handleNameChange(e, post.Id)}
                                                        value={post.Name}
                                                    />
                                                </div>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td colSpan="3">
                                                <div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
                                                    <label>Product price</label>
                                                    <br></br>
                                                    <input
                                                     style={{ width: "45rem" }}
                                                        readOnly={this.state.hiddenEditInfo}
                                                        className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
                                                        placeholder="Product price"
                                                        type="text"
                                                        onChange={e => this.handlePriceChange(e, post.Id)}
                                                        value={post.Price}
                                                    />
                                                </div>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td colSpan="3">
                                                <div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
                                                    <label>Product quantity</label>
                                                    <br></br>
                                                    <input
                                                     style={{ width: "45rem" }}
                                                        readOnly={this.state.hiddenEditInfo}
                                                        className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
                                                        placeholder="Product quantity"
                                                        type="text"
                                                        onChange={e => this.handleQuantityChange(e, post.Id)}
                                                        value={post.Quantity}
                                                    />
                                                </div>


                                            </td>
                                        </tr>

                                        <tr style={{ width: "100%" }}  >
                                            <td>
                                                <button onClick={() => this.AddToCart(post)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height: "40px", width: "6rem", marginLeft: "20rem" }}>Add to cart</button>
                                            </td>


                                        </tr>

                                        <br />
                                        <br />
                                        <br />
                                    </tr>

                                ))}

                            </tbody>
                        </table>
                    </div>
                </div>



                <Order
                    buttonName="Add"
                    header="Add product to cart"
                    show={this.state.showOrderModal}
                    onCloseModal={this.handleOrderModalClose}
                    handleQuantity={this.handleQuantityOrderChange}
                    product={this.state.product}
                />

            </React.Fragment>
        );
    }
}

export default ProfilePage;