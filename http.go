func (s) start() error{
	listen, err = net.Listen(s.network, s.address)
	if err != nil{
		return err
	}
	s.listen = listen
	if err := s.Serve(listen); !errors.Is(err, http.ErrServerClosed){
		return err
	}
	return nil
}

func (s) stop() error{
	return s.Shutdown()
}

