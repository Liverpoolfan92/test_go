﻿using Docker.DotNet;
using Docker.DotNet.Models;
using System;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Text.Json;

namespace ConsoleApp1
{
    class Program
    {
        static void Main(string[] args)
        {

            string hostName = Dns.GetHostName();
            IPAddress[] ipAddresses = Dns.GetHostAddresses(hostName);
            foreach (IPAddress ipAddress in ipAddresses)
            {
                if (ipAddress.AddressFamily == AddressFamily.InterNetwork)
                {
                    Console.WriteLine("IPv4 Address: " + ipAddress);
                }
            }

            int localPort = 8484;
            int targetPort = 8485;
            IPAddress localAddress = IPAddress.Parse("172.20.160.1");
            IPAddress targetAddress = IPAddress.Parse("172.17.0.1");

            try
            {
                var listener = new TcpListener(localAddress, localPort);
                listener.Start();

                Console.WriteLine($"Listening on {localAddress}:{localPort}");

                TcpClient client1 = listener.AcceptTcpClient();
                Console.WriteLine("Received connection");

                var networkStream = client1.GetStream();
                var buffer = new byte[1024];
                int bytesRead = networkStream.Read(buffer, 0, buffer.Length);

                var json = JsonSerializer.Deserialize<Interfaces>(Encoding.UTF8.GetString(buffer, 0, bytesRead));

                Console.WriteLine("Received interfaces:");
                for (int i = 0; i < json.InterfaceList.Length; i++)
                {
                    Console.WriteLine($"{i}: {json.InterfaceList[i]}");
                }

                Console.WriteLine("Choose an interface by its number:");
                int choice = int.Parse(Console.ReadLine());

                var chosenData = new Chosen { Interface = json.InterfaceList[choice] };
                var jsonBytes = JsonSerializer.SerializeToUtf8Bytes(chosenData);

                TcpClient targetClient = new TcpClient();
                targetClient.Connect(targetAddress, targetPort);
                Console.WriteLine($"Connected to {targetAddress}:{targetPort}");

                var targetStream = targetClient.GetStream();
                targetStream.Write(jsonBytes, 0, jsonBytes.Length);
                Console.WriteLine("Sent data");

                targetClient.Close();
                listener.Stop();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }

    }

    class Interfaces
    {
        public string[] InterfaceList { get; set; }
    }

    class Chosen
    {
        public string Interface { get; set; }
    }
}